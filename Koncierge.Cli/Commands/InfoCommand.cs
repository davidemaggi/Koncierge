using Spectre.Console;
using Spectre.Console.Cli;
using System.ComponentModel;

namespace Koncierge.Cli.Commands
{
    public class InfoCommand : Command<InfoCommand.Settings>
    {
        public class Settings : CommandSettings
        {

            //public string Name { get; set; }
        }

        public override int Execute(CommandContext context, Settings settings)
        {
            AnsiConsole.MarkupLine($"Hello [bold yellow]Koncierge[/]!");
            return 0;
        }
    }
}
