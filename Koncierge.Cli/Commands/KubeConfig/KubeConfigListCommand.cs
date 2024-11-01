

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Koncierge.Domain;
using Koncierge.Core.Services.Implementations;
using Koncierge.Core.Services.Interfaces;
using Spectre.Console;
using Spectre.Console.Cli;

namespace Koncierge.Cli.Commands.KubeConfig
{
    public class KubeConfigListCommand : Command<KubeConfigListCommand.Settings>
    {
        public class Settings : CommandSettings
        {

            //public string Name { get; set; }
        }

        private readonly IKonciergeService _konciergeService;

        public KubeConfigListCommand(IKonciergeService konciergeService)
        {
            _konciergeService = konciergeService ?? throw new ArgumentNullException(nameof(konciergeService));
        }

        public override int Execute(CommandContext context, Settings settings)
        {
            var kubeConfigs = _konciergeService.GetKubeConfigs().ToList();


            // Create a table
            var table = new Table();

            // Add some columns
            table.AddColumn("Name");
            table.AddColumn("Path");
            table.AddColumn(new TableColumn("IsDefault").Centered());

            kubeConfigs.ForEach(kubeConfig => table.AddRow(new Markup(kubeConfig.Name), new Markup(kubeConfig.Path), new Markup(kubeConfig.IsDefault ? Emoji.Known.CheckMarkButton : Emoji.Known.CrossMark).Centered()));

           
            


            // Render the table to the console
            AnsiConsole.Write(table);





            return 0;
        }
    }
}
