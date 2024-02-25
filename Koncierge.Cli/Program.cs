using Koncierge.Cli.Commands;
using Koncierge.Cli.Commands.KubeConfig;
using Koncierge.Cli.Infrastructure;
using Koncierge.Core;
using Koncierge.Database;
using Microsoft.Extensions.DependencyInjection;
using Spectre.Console;
using Spectre.Console.Cli;
using System;

namespace Koncierge.Cli
{
    internal class Program
    {
        static int Main(string[] args)
        {

            Console.OutputEncoding = System.Text.Encoding.UTF8;




            //var registrar = new TypeRegistrar(services);
            // Create a new command app with the registrar
            var services = new ServiceCollection();

            var app = new CommandApp<InfoCommand>(ConfigureServices(services));

            app.Configure(config =>
            {
                config.SetApplicationName("koncierge");

                config.AddBranch<KubeConfigSettings>("config", add =>
                {
                    add.AddCommand<GetKubeConfigCommand>("get").WithAlias("g").WithDescription("Search For kubeconfig files in a path, if you don't provide a path it will look into the .kube folder");
                    add.AddCommand<ListKubeConfigCommand>("list").WithAlias("l").WithAlias("ls").WithDescription("List of all the KubeConfig files Koncierge knows");
                    add.AddCommand<RemoveKubeConfigCommand>("remove").WithAlias("r").WithAlias("rm").WithDescription("Remove KubeConfig File(s) from Configuration");
                    add.AddCommand<MergeKubeConfigCommand>("merge").WithAlias("m").WithAlias("mrg").WithDescription("Merge 2 KubeConfig Files");
                    //add.AddCommand<AddReferenceCommand>("reference");
                }).WithAlias("c");

                

                config.AddCommand<InfoCommand>("info")
              .WithAlias("i")
              .WithAlias("v")
              .WithAlias("version")
              .WithDescription("Get info about the current configuration");


                config.SetExceptionHandler((ex) =>
                {

                    AnsiConsole.MarkupLine($":skull: {ex.Message}");


#if DEBUG
                    AnsiConsole.WriteException(ex, ExceptionFormats.ShortenEverything);
#endif
                });



#if DEBUG
                config.ValidateExamples();
#endif

            });



            try
            {
                app.Run(args);
                return 0;
            }
            catch (Exception ex)
            {
                AnsiConsole.MarkupLine($":skull: {ex.Message}");

                return -99;
            }
        }


        public static TypeRegistrar ConfigureServices(IServiceCollection services)
        {

            services.AddSingleton<IKonciergeDbService, KonciergeDbService>();
            services.AddSingleton<IKonciergeCoreService, KonciergeCoreService>();
            services.AddSingleton<HelpersService>();
            return new TypeRegistrar(services);



        }


    }
}
