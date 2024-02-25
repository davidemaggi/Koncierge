using Koncierge.Core;
using Koncierge.Models;
using Spectre.Console.Cli;
using Spectre.Console;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.ComponentModel;
using Koncierge.Exceptions;

namespace Koncierge.Cli.Commands.KubeConfig
{

    public class MergeKubeConfigSettings : KubeConfigSettings
    {
        [CommandOption("-s|--source")]
        public string inputSourcePath { get; set; }

        [CommandOption("-t|--target")]
        public string? inputTargetPath { get; set; }


    }
    public class MergeKubeConfigCommand : AsyncCommand<MergeKubeConfigSettings>
    {

        private readonly IKonciergeCoreService _konciergeCore;
        private readonly HelpersService _helper;
        public MergeKubeConfigCommand(IKonciergeCoreService kc, HelpersService hs)
        {
            _konciergeCore = kc;
            _helper = hs;
        }

        public override async Task<int> ExecuteAsync(CommandContext context, MergeKubeConfigSettings settings)
        {

           
            KubeConfigFile sourcetFile= new KubeConfigFile();
            KubeConfigFile targetFile = new KubeConfigFile();

            var list = new List<KubeConfigFile>();
            await AnsiConsole.Status()
   .StartAsync(":magnifying_glass_tilted_left: [bold]Getting the list of KubeConfig Files[/]", async ctx =>
   {
       list = await _konciergeCore.GetKubeConfigList();
   });

            if (settings.inputSourcePath is null) {
                throw MissingMandatoryParameterException.WithParameter("--source");
            }

            if (!await _konciergeCore.IsValidKubeConfig(settings.inputSourcePath))
            {
                throw KubeConfigNotValidException.WithKcf(settings.inputSourcePath);
            }
            sourcetFile = await _konciergeCore.GetKubeConfigFileFromPath(settings.inputSourcePath);


            if (settings.inputTargetPath is null)
            {

                _helper.WriteInfo("A target has not been provided, select it from the list of known KubeConfigs.");
                var strToMerge = new List<string>();
                foreach (var config in list)
                {

                    strToMerge.Add($"{config.Id} - {config.Name} @ {config.Path}");



                }


                var strSelect = AnsiConsole.Prompt(
    new SelectionPrompt<string>()
        .Title("Which KubeConfig you want to [green]MERGE[/]?")
        .PageSize(10)
        .MoreChoicesText("[grey](Move up and down to reveal more Configs)[/]")
        .AddChoices(strToMerge));

                _helper.WriteSelect(strSelect);

                targetFile = list.Where(x => x.Path == getIdFromSelection(strSelect)).FirstOrDefault();

            } 

        
        
            else {


                targetFile = list.Where(x => x.Path == settings.inputTargetPath).FirstOrDefault();

                if (targetFile is null)
                {
                    targetFile=await _konciergeCore.GetKubeConfigFileFromPath(settings.inputTargetPath);

                    if (targetFile is null)
                    {
                        throw KubeConfigNotFountException.WithKcf(settings.inputTargetPath);
                    }

                }



            }


           
        MergeResult merged = new MergeResult();

                 await AnsiConsole.Status()
   .StartAsync("[bold]Merging KubeConfig Files[/]", async ctx =>
   {
       merged = await _konciergeCore.MergeKubeConfig(sourcetFile.Path, targetFile.Path);
   });


            if (merged.DoneSomething()) { 

                var root = new Tree("");

                // Add some nodes
                var outAdded = root.AddNode("[green]Added[/]");

                foreach (var item in merged.Added)
                {
                    var tmp = outAdded.AddNode(item);
                }


                var outEdit = root.AddNode("[yellow]Modified[/]");
                foreach (var item in merged.Modified)
                {
                    var tmp = outEdit.AddNode(item);
                }
                var outDeleted = root.AddNode("[red]Deleted[/]");

                foreach (var item in merged.Deleted)
                {
                    var tmp = outDeleted.AddNode(item);
                }

                // Render the tree
                AnsiConsole.Write(root);






                if (AnsiConsole.Confirm("Confirm Merge?"))
                {
                    _konciergeCore.SaveKubeConfig(targetFile.Path, merged.Merged, true);
                }
                else
                {
                    _helper.WriteWarning("Merging Aborted");
                }




            _helper.WriteSuccess("Merging Completed");

            }else { 
            _helper.WriteInfo("Nothing to do in this merge.");

            }



            return 0;
        }


        public string getIdFromSelection(string str)
        {


            str = str.Replace(" ", "");
            var x = str.Split('@');

            return x[1];


        }

    }
}
