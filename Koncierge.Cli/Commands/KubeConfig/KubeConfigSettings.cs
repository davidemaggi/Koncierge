using Spectre.Console.Cli;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Cli.Commands.KubeConfig
{
    public class KubeConfigSettings : GlobalSettings
    {
        [CommandOption("-d|--dry")]
        [DefaultValue(false)]
        public bool DryRun { get; set; }
    }
}
