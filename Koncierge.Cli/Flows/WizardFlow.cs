using Koncierge.Cli.Commands.Printers;
using Koncierge.Cli.Components;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using Koncierge.Domain.Enums;
using Spectre.Console;

namespace Koncierge.Cli.Flows;

public static class WizardFlow
{
    private static  IKonciergeService _konciergeService;

    public static Task<int> RunAsync(IKonciergeService konciergeService, bool isDry)
    {
        _konciergeService = konciergeService;
        AnsiConsole.WriteLine($"{Emoji.Known.Mage} Port Forward Wizard");

            var tmpFwd = new ForwardEntity();

            #region KubeConfig

            

          
            var configs = _konciergeService.GetKubeConfigs();
            var conf = MultiselectComponent<KubeConfigEntity>.One(configs, "KubeConfig");

            if (conf==default!)
            {
                AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No [{Color.Red}]KubeConfig[/] have been found");
                return Task.FromResult(1);
            }
            
            AnsiConsole.MarkupLine($"{Emoji.Known.BookmarkTabs} KubeConfig: [{Color.DeepSkyBlue1}]{conf}[/]");

            tmpFwd.WithConfig = conf;

            #endregion
            
            #region Context
            
            var contexts = _konciergeService.GetContextsForConfig(conf.Id).AsQueryable();

            var selContext = MultiselectComponent<KubeContextDto>.One(contexts, "Context");

            if (selContext==default!)
            {
                AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No [{Color.Red}]Context[/] have been found");
                return Task.FromResult(1);
            }
            
            

            AnsiConsole.MarkupLine($"{Emoji.Known.PageFacingUp} Context: [{Color.DeepSkyBlue1}]{selContext}[/]");
            tmpFwd.Context = selContext.Name;

            #endregion
            
            var idConn = _konciergeService.ConnectTo(conf, selContext.Name);

            
            #region Namespace
            var namespaces = _konciergeService.GetNamespacesForConnection(idConn).AsQueryable();

            
            var selNamespace = MultiselectComponent<KubeNamespaceDto>.One(namespaces, "Namespace");

            if (selNamespace==default!)
            {
                AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No [{Color.Red}]Namespace[/] have been found");
                return Task.FromResult(1);
            }
            

           

            AnsiConsole.MarkupLine($"{Emoji.Known.PeopleHugging} NameSpace: [{Color.DeepSkyBlue1}]{selNamespace}[/]");

            tmpFwd.Namespace = selNamespace.Name;


#endregion

#region serviceorPod
            var selfwdType = AnsiConsole.Prompt(
                new SelectionPrompt<string>()
                    .Title("Which [green]kind[/] of entity you want to forward?")
                    //.PageSize(10)
                    //.MoreChoicesText("[grey](Move up and down to reveal more namespace)[/]")
                    .AddChoices(new[] { "Pod", "Service" })

                    .EnableSearch()
            );

            var fwdType = selfwdType.ToLower().Equals("service")
                ? KonciergeForwardType.Service
                : KonciergeForwardType.Pod;


            AnsiConsole.MarkupLine(
                $"{Emoji.Known.Compass} Type: {(fwdType == KonciergeForwardType.Service ? Emoji.Known.GlobeWithMeridians : Emoji.Known.Package)} [{Color.DeepSkyBlue1}]{selfwdType}[/]");


            tmpFwd.Type = fwdType;
            
            
            KubeForwardableDto forward = default!;
            if (tmpFwd.Type == KonciergeForwardType.Service)
            {

                var services = _konciergeService.GetServicesForNamespace(idConn, tmpFwd.Namespace).AsQueryable();
                
                forward = MultiselectComponent<KubeForwardableDto>.One(services, "Services");

                if (selNamespace==default!)
                {
                    AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No forwardable [{Color.Red}]service[/] have been found");
                    return Task.FromResult(1);
                }
                
                
                
                
            }

            if (tmpFwd.Type == KonciergeForwardType.Pod)
            {

                var pods = _konciergeService.GetPodsForNamespace(idConn, tmpFwd.Namespace).AsQueryable();



                forward = MultiselectComponent<KubeForwardableDto>.One(pods, "Pods");

                if (selNamespace==default!)
                {
                    AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No forwardable [{Color.Red}]service[/] have been found");
                    return Task.FromResult(1);
                }


            }

            if (forward == default!)
            {
                AnsiConsole.MarkupLine(
                    $"{Emoji.Known.StopSign} No [{Color.Red}]Forwardable entity[/] have been found");
                return Task.FromResult(1);

            }

            AnsiConsole.MarkupLine(
                $"{Emoji.Known.Label} Forwarding: {(tmpFwd.Type == KonciergeForwardType.Service ? Emoji.Known.GlobeWithMeridians : Emoji.Known.Package)} [{Color.DeepSkyBlue1}]{forward.Name}[/]");

            tmpFwd.Selector = forward.Name;
            #endregion

                
                
               

#region port
            var selPort = MultiselectComponent<KubePortDto>.One(forward.Ports.AsQueryable(), "Port");

            if (selPort==default!)
            {
                AnsiConsole.MarkupLine($"{Emoji.Known.StopSign} No forwardable [{Color.Red}]port[/] have been found");
                return Task.FromResult(1);
            }
                
                
                    
                    
            AnsiConsole.MarkupLine($"{Emoji.Known.Door} Remote Port: [{Color.DeepSkyBlue1}]{selPort.ToString()}[/]");

            tmpFwd.RemotePort = selPort.Port;
            var validLocalPort = false;
            var localPort = 0;
                    
            while (!validLocalPort)
            {
                localPort = AnsiConsole.Prompt(
                    new TextPrompt<int>($"{Emoji.Known.InputNumbers} On Which Local Port?"));
                        
                // TODO: Validare LocalPort
                validLocalPort=localPort >=0 && localPort<=65535;

                if (!validLocalPort)
                {
                    AnsiConsole.MarkupLine(
                        $"{Emoji.Known.StopSign} Local port {localPort} is [{Color.Red}]invalid[/]");
                }

            }


            var shouldSave = ConfirmationComponent.Ask("Do you want to save this forward configuration?");

            tmpFwd.LocalPort = localPort;

                    //AnsiConsole.MarkupLine($"{Emoji.Known.DesktopComputer} You can use the command: {_konciergeService.buildForwardCommandString(tmpFwd)}");
                  
                    

#endregion


#region secretsConfig
KubeConfigDto mapsAndSecrets = new KubeConfigDto();
tmpFwd.Configs=new List<LinkedConfig>();
if (shouldSave)
{
    var addConf = ConfirmationComponent.Ask("Do you want to link one or more Secrets/ConfigMap?");
    if (addConf)
    {

        var more = true;
        mapsAndSecrets = _konciergeService.GetKubeMapAndSecretsForNamespace(idConn, selNamespace.Name);

        while (more && (mapsAndSecrets.Secrets.Any() || mapsAndSecrets.ConfigMaps.Any()))
        {
            
            
            var confType = MultiselectComponent<string>.One( (new string[]{"🔑 Secret", "⚙️ ConfigMap"}).AsQueryable(), "Type");

            var isSecret = confType == "🔑 Secret";
            
            var main =isSecret? MultiselectComponent<ConfigItemDto>.One( mapsAndSecrets.Secrets.AsQueryable(), "Secret"): MultiselectComponent<ConfigItemDto>.One( mapsAndSecrets.ConfigMaps.AsQueryable(), "ConfigMap");

            var inners = MultiselectComponent<string>.Many(main.Values.Select(x=>x.Key).AsQueryable(), "Which one", 10, false);

            if (tmpFwd.Configs is null)
            {
                
            }

            var confToAdd = tmpFwd.Configs.FirstOrDefault(x => x.Name == main.Selector);
            
            if (confToAdd is null)
            {
                             confToAdd = new LinkedConfig();
                             confToAdd.Name = main.Selector;
                             confToAdd.Type = isSecret ? KonciergeConfigType.Secret : KonciergeConfigType.ConfigMap;
                             confToAdd.Values = new List<string>();
            }
            else
            {

                if (confToAdd.Values is null)
                {
                    confToAdd.Values = new List<string>();


                }  
            }

            foreach (var inner in inners)
            {
                if (!confToAdd.Values.Contains(inner))
                {
                    confToAdd.Values.Add(inner);
                }
            }
            
            

            tmpFwd.Configs.Add(confToAdd);
            
            

            
            more = ConfirmationComponent.Ask("Another one?");

            
        }



        _konciergeService.SaveForward(tmpFwd);



    }

}
else
{
    AnsiConsole.MarkupLine($"{Emoji.Known.Droplet} This is a dry run, the configuration has not been saved");
}
#endregion


            if (!isDry) {
                AnsiConsole.Clear();
                RunningFwdPrinter.Print(tmpFwd,mapsAndSecrets);

                _konciergeService.ExecuteKubectl(_konciergeService.buildForwardCommandString(tmpFwd)).Wait();
            }
            return Task.FromResult(0);



        
        
    }



}