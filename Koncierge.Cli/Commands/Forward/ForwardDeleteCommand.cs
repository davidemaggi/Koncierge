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
using System.ComponentModel;

namespace Koncierge.Cli.Commands.Forward
{
    internal class ForwardDeleteCommand : AsyncCommand<ForwardDeleteCommand.Settings>
    {
        public class Settings : CommandSettings
        {

            [CommandOption("--all")]
            [DefaultValue(false)]
            public bool All { get; set; }
        }


        private readonly IKonciergeService _konciergeService;
        private readonly IKubeForwardRepository _kubeForwardRepository;


        public ForwardDeleteCommand(IKonciergeService konciergeService, IKubeForwardRepository kubeForwardRepository)
        {
            _konciergeService = konciergeService ?? throw new ArgumentNullException(nameof(konciergeService));
            _kubeForwardRepository = kubeForwardRepository ?? throw new ArgumentNullException(nameof(kubeForwardRepository));
        }



        public override Task<int> ExecuteAsync(CommandContext context, Settings settings)
        {







            var knownForward = _kubeForwardRepository.GetAllWithInclude().OrderBy(x => x.WithConfig.Name).ThenBy(x => x.Context).ThenBy(x => x.Namespace).ThenBy(x => x.Type).ThenBy(x => x.Selector).ThenBy(x => x.LocalPort).ToList();


            var xxx = knownForward.GroupBy(x => new { x.WithConfig.Id, x.WithConfig.Name, x.Context }).ToList();




            List<ForwardEntity> selFwd = new List<ForwardEntity>();

            if (settings.All)
            {
                selFwd.AddRange(knownForward);

                AnsiConsole.MarkupLine($"{Emoji.Known.Axe} Delete [{Color.Red}]All[/] forwards");


            }
            else
            {
                if (knownForward.Count() > 1)
                {

                    var prm = new MultiSelectionPrompt<ForwardEntity>()
                            .Title("Select the [red]Forward[/] you want to delete.")
                            .PageSize(10)

                            .MoreChoicesText("[grey](Move up and down to reveal more forwards)[/]")
                            ;

                    foreach (var x in xxx)
                    {
                        prm.AddChoiceGroup<ForwardEntity>(new ForwardEntity() { Id = Guid.Empty, Context = $"{x.Key.Name} -> {x.Key.Context}" }, x);

                    }

                    selFwd = AnsiConsole.Prompt(
                       prm
                    );

                }
               
                else
                {

                    AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No [{Color.Red}]Forward[/] has been found");
                    return Task.FromResult(1);

                }
            }

            foreach (var forwardEntity in selFwd.Where(x => x.Id != Guid.Empty))
            {
                AnsiConsole.MarkupLine($"{Emoji.Known.Axe} Selected: {forwardEntity.Selector}:[{Color.DeepSkyBlue1}]{forwardEntity.RemotePort}[/] -> localhost:[{Color.DeepSkyBlue1}]{forwardEntity.LocalPort}[/]");


            }


            var sure = AnsiConsole.Prompt(
               new TextPrompt<bool>("Are you really sure to delete these Forward configuration?")
                   .AddChoice(true)
                   .AddChoice(false)
                   .DefaultValue(true)
                   .WithConverter(choice => choice ? "y" : "n"));

            if (sure)
            {
                foreach (var forwardEntity in selFwd.Where(x => x.Id != Guid.Empty))
                {
                    AnsiConsole.MarkupLine($"{Emoji.Known.Axe} Deleting: {forwardEntity.Selector}:[{Color.DeepSkyBlue1}]{forwardEntity.RemotePort}[/] -> localhost:[{Color.DeepSkyBlue1}]{forwardEntity.LocalPort}[/]");


                    _kubeForwardRepository.Delete(forwardEntity.Id);

                }

            }
            else {

                AnsiConsole.MarkupLine($"{Emoji.Known.BeerMug} That's [{Color.DeepSkyBlue1}]OK[/] we can still be friend");


            }



            return Task.FromResult(0);
        }


    }
}
