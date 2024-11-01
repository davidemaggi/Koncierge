using k8s.KubeConfigModels;
using Koncierge.Cli.Commands.Tools;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain;
using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using Koncierge.Domain.Enums;
using Koncierge.Domain.Repositories.Interfaces;
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

namespace Koncierge.Cli.Commands.Forward
{
    internal class StartCommand : AsyncCommand<StartCommand.Settings>
    {
        public class Settings : CommandSettings
        {

            [CommandOption("--all")]
            [DefaultValue(false)]
            public bool All { get; set; }
        }


        private readonly IKonciergeService _konciergeService;
        private readonly IKubeForwardRepository _kubeForwardRepository;


        public StartCommand(IKonciergeService konciergeService, IKubeForwardRepository kubeForwardRepository)
        {
            _konciergeService = konciergeService ?? throw new ArgumentNullException(nameof(konciergeService));
            _kubeForwardRepository= kubeForwardRepository ?? throw new ArgumentNullException(nameof(kubeForwardRepository));
        }



        public override Task<int> ExecuteAsync(CommandContext context, Settings settings)
        {

return StartFlow.RunAsync(_konciergeService, settings.All);
            
        }

       
    }
}
