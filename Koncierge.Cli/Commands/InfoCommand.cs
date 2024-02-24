using Spectre.Console.Cli;
using Spectre.Console;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Cli.Commands
{
    public class InfoCommandSettings : GlobalSettings { }

    public class InfoCommand : AsyncCommand<InfoCommandSettings>
    {

         public class InfoCommandSettings : GlobalSettings { }
        public InfoCommand()
        {

        }

        public override async Task<int> ExecuteAsync(CommandContext context, Commands.InfoCommandSettings settings)
        {
            AnsiConsole.Write(
    new FigletText("Koncierge")
        .LeftJustified()
        .Color(Color.DarkCyan));

            Emoji.Remap("big_hearth", "❤️");

            var rule = new Rule("Made with :big_hearth: in [default on green] [/][default on white] [/][default on red] [/]");
            rule.Centered();

            AnsiConsole.Write(rule);


            return 0;
        }

        public string GetAssemblyVersion()
        {
            return GetType().Assembly.GetName().Version.ToString();
        }

        public string GetAssemblyVersionInfo()
        {
            var assembly = Assembly.GetExecutingAssembly();
            var informationVersion = assembly.GetCustomAttribute<AssemblyInformationalVersionAttribute>().InformationalVersion;
            return informationVersion;
        }

        
    }
}