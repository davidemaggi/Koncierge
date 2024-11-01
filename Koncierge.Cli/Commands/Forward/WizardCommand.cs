 using Koncierge.Cli.Commands.Tools;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain;
using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using Spectre.Console;
using Spectre.Console.Cli;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Koncierge.Cli.Components;
using Koncierge.Cli.Flows;
using Koncierge.Domain.Enums;

namespace Koncierge.Cli.Commands.Forward
{
    internal class WizardCommand : AsyncCommand<WizardCommand.Settings>
    {
        public class Settings : CommandSettings
        {

            [CommandOption("--dry")]
            [DefaultValue(false)]
            public bool IsDry { get; set; }
        }


        private readonly IKonciergeService _konciergeService;

        public WizardCommand(IKonciergeService konciergeService)
        {
            _konciergeService = konciergeService ?? throw new ArgumentNullException(nameof(konciergeService));
        }



        public override Task<int> ExecuteAsync(CommandContext context, Settings settings)
        {





            return WizardFlow.RunAsync(_konciergeService, settings.IsDry);
        }

    }
} 
