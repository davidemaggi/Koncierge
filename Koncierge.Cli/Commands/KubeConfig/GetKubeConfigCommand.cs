using Koncierge.Cli.Commands.KubeConfig;
using Koncierge.Core;
using Koncierge.Models.Config;
using Spectre.Console;
using Spectre.Console.Cli;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Cli.Commands.KubeConfig
{
    public class GetKubeConfigSettings : KubeConfigSettings {

        [CommandOption("-p|--path")]
        public string? inputPath { get; set; }

    }
    public class GetKubeConfigCommand : AsyncCommand<GetKubeConfigSettings>
    {

        private readonly IKonciergeCoreService _konciergeCore;
        public GetKubeConfigCommand(IKonciergeCoreService kc)
        {
            _konciergeCore = kc;
        }

        public override async Task<int> ExecuteAsync(CommandContext context, GetKubeConfigSettings settings)
        {

            
        var list = new List<KubeConfigFile>();
            await AnsiConsole.Status()
   .StartAsync(":magnifying_glass_tilted_left: [bold]Searching Kube Config Files[/]", async ctx =>
   {
   if (settings.inputPath is null)
        {
               
        list = await _konciergeCore.GetKubeConfigFromPath();
 
                
        }
        else
        {
                
        list = await _konciergeCore.GetKubeConfigFromPath(settings.inputPath);
  
                
        }
      });
                if (list.Count > 0)
        {
            AnsiConsole.MarkupLine($":party_popper: [bold]{list.Count}[/] files have been found");
        }
        else
        {
            AnsiConsole.MarkupLine($":crying_face: no files have been found");
        }

        
        var table = new Table();

        // Add some columns
        table.AddColumn(new TableColumn("Status").Centered());
        table.AddColumn(new TableColumn("Name").Centered());
        table.AddColumn(new TableColumn("Path").Centered());



        foreach (var item in list)
        {

            table.AddRow(item.Status.ToString(), item.Name, item.Path);


        }

            AnsiConsole.MarkupLine("");

            AnsiConsole.Write(table);
   

            
            return 0;
        }




    }
}
