using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain.Entities;
using Koncierge.Domain.Repositories.Interfaces;
using Spectre.Console.Cli;
using Spectre.Console;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Koncierge.Domain.Enums;

namespace Koncierge.Cli.Commands.Forward
{
    internal class ForwardListCommand : AsyncCommand<ForwardListCommand.Settings>
    {
        public class Settings : CommandSettings
        {

            //public string Name { get; set; }
        }


        private readonly IKonciergeService _konciergeService;
        private readonly IKubeForwardRepository _kubeForwardRepository;


        public ForwardListCommand(IKonciergeService konciergeService, IKubeForwardRepository kubeForwardRepository)
        {
            _konciergeService = konciergeService ?? throw new ArgumentNullException(nameof(konciergeService));
            _kubeForwardRepository = kubeForwardRepository ?? throw new ArgumentNullException(nameof(kubeForwardRepository));
        }



        public override Task<int> ExecuteAsync(CommandContext context, Settings settings)
        {


           // AnsiConsole.WriteLine($"{Emoji.Known.MagnifyingGlassTiltedLeft} Start forwarding an existing configuration");





            var knownForward = _kubeForwardRepository.GetAllWithInclude().OrderBy(x=>x.WithConfig.Name).ThenBy(x=>x.Context).ThenBy(x => x.Namespace).ThenBy(x => x.Type).ThenBy(x => x.Selector).ThenBy(x => x.LocalPort).ToList();


            var table = new Table();

            // Add some columns
            table.AddColumn("Config");
            table.AddColumn("Context");
            table.AddColumn("Namespace");
            table.AddColumn("Type");
            table.AddColumn("Selector");
            table.AddColumn("Port");
           

            knownForward.ForEach(fwd => table.AddRow(new Markup(fwd.WithConfig.Name), new Markup(fwd.Context), new Markup(fwd.Namespace), new Markup(fwd.Type == KonciergeForwardType.Service ? Emoji.Known.GlobeWithMeridians : Emoji.Known.Package).Centered(), new Markup(fwd.Selector), new Markup($"{fwd.LocalPort}:{fwd.RemotePort}")));





            // Render the table to the console
            AnsiConsole.Write(table);







            
            return Task.FromResult(0);
        }


    }
}
