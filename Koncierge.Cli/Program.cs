using Koncierge.Cli.Commands;
using Koncierge.Cli.Commands.Forward;
using Koncierge.Cli.Commands.KubeConfig;
using Koncierge.Cli.Commands.Tools;
using Koncierge.Cli.Injection;
using Koncierge.Core.Services.Implementations;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain;
using Koncierge.Domain.Repositories.Implementations;
using Koncierge.Domain.Repositories.Interfaces;
using Microsoft.Extensions.DependencyInjection;
using Spectre.Console.Cli;

namespace Koncierge.Cli
{
    internal class Program
    {
        public static int Main(string[] args)
        {
            var app = new CommandApp(RegisterServices());
            app.Configure(config => ConfigureCommands(config));

            return app.Run(args);
        }




        private static IConfigurator ConfigureCommands(IConfigurator config)
        {
            config.CaseSensitivity(CaseSensitivity.None);
            config.SetApplicationName("Koncierge");
            config.ValidateExamples();


            config.AddCommand<InfoCommand>("info")
               .WithAlias("i")
                   .WithDescription("Get Info About Koncierge.")
                   // .WithExample(new[] { "hello", "--name", "DarthPedro" })
                   ;

           


            config.AddBranch("kubeconfig", kc =>
            {
                kc.SetDescription("Manage KubeConfigs.");

                kc.AddCommand<KubeConfigListCommand>("list")
                    .WithAlias("ls")
                    .WithDescription("Get list of known KubeConfig.")
                    .WithExample(new[] { "kubeconfig", "ls" })
                    .WithExample(new[] { "kubeconfig", "list" })
                    ;

               
            });

            config.AddBranch("forward", fwd =>
            {
                
                fwd.SetDescription("Forward Actions");
                fwd.SetDefaultCommand<WizardCommand>();

                fwd.AddCommand<WizardCommand>("wizard")
                    .WithAlias("w")
                    .WithDescription("Run Wizard.")
                    .WithExample(new[] { "fwd" })
                    .WithExample(new[] { "forward" })
                    ;

                fwd.AddCommand<StartCommand>("start")
                    .WithAlias("s")
                    .WithDescription("Start Fowward")
                    .WithExample(new[] { "fwd","start" })
                    .WithExample(new[] { "forward","start" })
                    ;

                fwd.AddCommand<ForwardListCommand>("list")
                    .WithAlias("ls")
                    .WithDescription("List all previously saved forwards")
                    .WithExample(new[] { "fwd", "ls" })
                    .WithExample(new[] { "forward", "list" })
                    ;

                fwd.AddCommand<ForwardDeleteCommand>("remove")
                   .WithAlias("rm")
                   .WithDescription("Delete previously saved forwards")
                   .WithExample(new[] { "fwd", "rm" })
                   .WithExample(new[] { "forward", "remove" })
                   ;


            }).WithAlias("fwd");

            /*
            config.AddBranch("tools", student =>
            {
                student.SetDescription("Generic Tools.");

                student.AddCommand<InitCommand>("init")
                    //.WithAlias("add")
                    .WithDescription("Initialize Koncierge.")
                    .WithExample(new[] { "tools", "init" });


            });
            */

            return config;
        }

        private static ITypeRegistrar RegisterServices()
        {
            // Create a type registrar and register any dependencies.
            // A type registrar is an adapter for a DI framework.
            var services = new ServiceCollection();

            // register services here...
            services.AddDbContext<KonciergeContext>();
            services.AddLogging();

            services.AddSingleton<IKubeConfigRepository, KubeConfigRepository>();
            services.AddSingleton<IKubeForwardRepository, KubeCForwardRepository>();
            services.AddSingleton<IKonciergeService, KonciergeService>();

            return new TypeRegistrar(services);
        }

    }
}
