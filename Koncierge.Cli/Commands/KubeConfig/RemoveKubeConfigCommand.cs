
using Koncierge.Cli.Commands.KubeConfig;
using Koncierge.Core;
using Koncierge.Models;
using Spectre.Console;
using Spectre.Console.Cli;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Cli.Commands.KubeConfig
{
    public class RemoveKubeConfigSettings : KubeConfigSettings {

        [CommandOption("-p|--path")]
        public string? inputPath { get; set; }


        [CommandOption("-i|--id")]
        public int? inputId { get; set; }

        [CommandOption("-f|--force")]
        [DefaultValue(false)]
        public bool inputForce { get; set; }

    }
    public class RemoveKubeConfigCommand : AsyncCommand<RemoveKubeConfigSettings>
    {
        private readonly HelpersService _helper;
        private readonly IKonciergeCoreService _konciergeCore;
        public RemoveKubeConfigCommand(IKonciergeCoreService kc, HelpersService hs)
        {
            _konciergeCore = kc;
            _helper = hs;
        }

        public override async Task<int> ExecuteAsync(CommandContext context, RemoveKubeConfigSettings settings)
        {

            var toDelete=new List<KubeConfigFile>();
            var strSelect = new List<string>();
            var strToDelete = new List<string>();
            await AnsiConsole.Status()
    .StartAsync("Getting Configs", async ctx =>
    {
        toDelete = await _konciergeCore.GetKubeConfigList();
    });
            

            if (settings.inputPath is not null) {

                var sel = toDelete.Where(x => x.Path == settings.inputPath).First();

                if (sel != null)
                {
                    strSelect.Add($"{sel.Id} - {sel.Name} @ {sel.Path}");
                }





            } else if(settings.inputId is not null) {
                var sel = toDelete.Where(x => x.Id == settings.inputId).FirstOrDefault();

                if (sel is not null)
                {
                    strSelect.Add($"{sel.Id} - {sel.Name} @ {sel.Path}");
                }
            } else {

                foreach (var config in toDelete) {

                    strToDelete.Add($"{config.Id} - {config.Name} @ {config.Path}");



                }


                strSelect = AnsiConsole.Prompt(
    new MultiSelectionPrompt<string>()
        .Title("Which KubeConfig you want [red]REMOVE[/]?")
        .NotRequired()
        .PageSize(10)
        .MoreChoicesText("[grey](Move up and down to reveal more Configs)[/]")
        .InstructionsText(
            "[grey](Press [blue]<space>[/] to toggle a Config, " +
            "[green]<enter>[/] to accept)[/]")
        .AddChoices(strToDelete));



            }


            if (strSelect.Count == 0)
            {
                _helper.WriteError("[red]No selection has been made or it has not been found[/]");
                
                return 0;
            }
            else {

                _helper.WriteSelect(strSelect);


                var rows = new List<Text>();

                foreach (var s in strSelect) { 
                
                    rows.Add(new Text(s));

                }

                AnsiConsole.Write(new Rows(rows));


                if (AnsiConsole.Confirm("Are you sure you want to proceed with the [red]REMOVAL[/]"))
                {

                    foreach (var s in strSelect)
                    {

                        await AnsiConsole.Status()
    .StartAsync("Deleting", async ctx =>
    {
        var tm = await _konciergeCore.RemoveKubeConfig(getIdFromSelection(s));
    });

                    }

                    
                    



                }
                else { 
               
                    _helper.WriteOk("Thats OK, we are still friends.");

                }




            }



            return 0;
        }


        public int getIdFromSelection(string str) {


            str = str.Replace(" ", "");
            var x = str.Split('-');

            return Int32.Parse(x[0]);
        
        
        }


    }
}
