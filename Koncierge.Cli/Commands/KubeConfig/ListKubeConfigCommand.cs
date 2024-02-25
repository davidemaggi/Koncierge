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
    public class ListKubeConfigSettings : KubeConfigSettings
    {



    }
    public class ListKubeConfigCommand : AsyncCommand<ListKubeConfigSettings>
    {

        private readonly IKonciergeCoreService _konciergeCore;
        public ListKubeConfigCommand(IKonciergeCoreService kc)
        {
            _konciergeCore = kc;
        }

        public override async Task<int> ExecuteAsync(CommandContext context, ListKubeConfigSettings settings)
        {




            var list = new List<KubeConfigFile>();
            await AnsiConsole.Status()
   .StartAsync(":magnifying_glass_tilted_left: [bold]Getting the list of KubeConfig Files[/]", async ctx =>
   {
       list = await _konciergeCore.GetKubeConfigList();
   });

            

            var table = new Table();
            table.Border = TableBorder.Ascii;

            // Add some columns
            table.AddColumn(new TableColumn("Id").Centered());
            table.AddColumn(new TableColumn("Status").Centered());
            table.AddColumn(new TableColumn("Name").Centered());
            table.AddColumn(new TableColumn("Path").Centered());



            foreach (var item in list)
            {
                table.AddRow(item.Id.ToString(), item.Status.ToString(), item.Name, item.Path);

            }


            AnsiConsole.MarkupLine($"Koncierge is aware of these {list.Count} KubeConfig");
            AnsiConsole.MarkupLine("");

            AnsiConsole.Write(table);



            return 0;
        }




    }
}
