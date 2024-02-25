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

            KubeConfigFile sourceFile;
            KubeConfigFile targetFile;

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

                targetFile = list.Where(x => x.Path == strSelect).FirstOrDefault();

            } 

        
        
            else {


                targetFile = list.Where(x => x.Path == settings.inputTargetPath).FirstOrDefault();

                if (targetFile is null)
                {
                    throw KubeConfigNotFountException.WithKcf(settings.inputTargetPath);
                }



            }


           



            



            return 0;
        }




    }
}
