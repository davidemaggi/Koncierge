using Koncierge.Cli.Commands.Printers;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain.Entities;
using Spectre.Console;

namespace Koncierge.Cli.Flows;

public static class StartFlow
{
    private static  IKonciergeService _konciergeService;

    public static Task<int> RunAsync(IKonciergeService konciergeService, bool isAll)
    {
        _konciergeService= konciergeService;
        AnsiConsole.WriteLine($"{Emoji.Known.MagnifyingGlassTiltedLeft} Start forwarding an existing configuration");





        //var knownForward = _konciergeService.GetAllForwards().OrderBy(x => x.WithConfig.Name).ThenBy(x => x.Context).ThenBy(x => x.Namespace).ThenBy(x => x.Type).ThenBy(x => x.Selector).ThenBy(x => x.LocalPort).ToList();
        var knownForward = _konciergeService.GetAllForwards().ToList();


           var grouped = knownForward.GroupBy(x => new { x.WithConfig.Id, x.WithConfig.Name, x.Context }).ToList();



            List<ForwardEntity> selFwd = new List<ForwardEntity>();

            if (isAll) {
                selFwd.AddRange(knownForward);

                AnsiConsole.MarkupLine($"{Emoji.Known.PartyPopper} Starting [{Color.DeepSkyBlue1}]All[/] forwards");


            }
            else { 
            if (knownForward.Count() > 1)
            {

                var prm = new MultiSelectionPrompt<ForwardEntity>()
                        .Title("Select the [green]Connection[/] you want to open.")
                        .PageSize(10)

                        .MoreChoicesText("[grey](Move up and down to reveal more forwards)[/]")
                        ;

                foreach (var x in grouped) {
                    prm.AddChoiceGroup<ForwardEntity>(new ForwardEntity() { Id = Guid.Empty, Context = $"{x.Key.Name} -> {x.Key.Context}" }, x);

                }

                selFwd = AnsiConsole.Prompt(
                   prm
                );

            }
            else if (knownForward.Count() == 1)
            {
                selFwd.AddRange(knownForward);

            }
            else
            {

                AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No [{Color.Red}]Forward[/] has been found");
                return Task.FromResult(1);

            }
        }
            var allTasks = new List<Task>();
            foreach (var forwardEntity in selFwd.Where(x=>x.Id!=Guid.Empty))
            {
                var idConn = _konciergeService.ConnectTo(forwardEntity.WithConfig, forwardEntity.Context);

                var mapsAndSecrets = _konciergeService.GetKubeMapAndSecretsForNamespace(idConn, forwardEntity.Namespace);

                
                
                
RunningFwdPrinter.Print(forwardEntity, mapsAndSecrets);
                allTasks.Add(_konciergeService.ExecuteKubectl(_konciergeService.buildForwardCommandString(forwardEntity)));
                
                
                
            }

            try
            {
                Task.WhenAll(allTasks).Wait();
            }
            catch (Exception e)
            {
                
                AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} Error [{Color.Red}]Forwarding[/] the required service(s)");

                
            }



            return Task.FromResult(0);
        
        
    }
}