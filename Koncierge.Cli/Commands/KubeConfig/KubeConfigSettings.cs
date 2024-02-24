using Spectre.Console.Cli;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Cli.Commands.KubeConfig
{
    public class KubeConfigSettings : GlobalSettings
    {
        [CommandOption("--dry")]
        public bool? DryRun { get; set; }
    }
}
