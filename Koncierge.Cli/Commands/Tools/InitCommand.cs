using Spectre.Console.Cli;
using Spectre.Console;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Koncierge.Domain;
using Koncierge.Core.Services.Interfaces;

namespace Koncierge.Cli.Commands.Tools
{
    public class InitCommand : Command<InitCommand.Settings>
    {
        public class Settings : CommandSettings
        {

            //public string Name { get; set; }
        }


        private readonly KonciergeContext _ctx;

        public InitCommand(KonciergeContext ctx)
        {
            _ctx = ctx ?? throw new ArgumentNullException(nameof(ctx));
        }



        public override int Execute(CommandContext context, Settings settings)
        {
            _ctx.Initialize();
            return 0;
        }
    }
}
