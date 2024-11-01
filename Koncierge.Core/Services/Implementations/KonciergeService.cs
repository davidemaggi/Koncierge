using k8s;
using k8s.KubeConfigModels;
using Koncierge.Core.Exceptions;
using Koncierge.Core.Models;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using Koncierge.Domain.Repositories.Interfaces;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Koncierge.Domain.Enums;
using Microsoft.Extensions.Logging;
using System.Net;
using Koncierge.Core.Tools;

namespace Koncierge.Core.Services.Implementations
{
    public class KonciergeService : IKonciergeService
    {
        private readonly List<KonciergeConnectionModel> _connections;
        private readonly IKubeConfigRepository _kubeConfigRepository;
        private readonly IKubeForwardRepository _kubeForwardRepository;

        private string _kubectl = "";
        private string _osName = "";

        private ILogger<KonciergeService> _logger { get; }



        public KonciergeService(ILogger<KonciergeService> logger , IKubeConfigRepository kubeConfigRepository, IKubeForwardRepository kubeForwardRepository) {
            _logger = logger;
            _kubeConfigRepository= kubeConfigRepository;
            _kubeForwardRepository = kubeForwardRepository;


            _connections =new List<KonciergeConnectionModel>();


            // Initialize


            Initialize();

        }


        private void Initialize() {

            var os = Environment.OSVersion;

            if (os.Platform == PlatformID.Win32NT) { _osName = $"windows"; _kubectl = "kubectl.exe"; }
            if (os.Platform == PlatformID.Unix) { _osName = $"linux"; _kubectl = "kubectl"; }
            if (os.Platform == PlatformID.Other) { _osName = $"darwin"; _kubectl = "kubectl"; }


            CheckAndAddDefaultConfig().Wait();

            CheckAndDownloadKubectl().Wait();

        }

        public Task ExecuteKubectl(string command) {

            return Task.Run(() =>
             {

                 
                 
                 
                 System.Diagnostics.Process process = new System.Diagnostics.Process();
                 System.Diagnostics.ProcessStartInfo startInfo = new System.Diagnostics.ProcessStartInfo();
                 startInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Normal;
                 startInfo.FileName = Path.Combine(FileSystemTools.getKonciergePath(), _kubectl);
                 startInfo.WorkingDirectory = FileSystemTools.getKonciergePath();
                 startInfo.Arguments = $"{command}";

                 startInfo.RedirectStandardOutput = true;
                 startInfo.RedirectStandardInput = true;
                 startInfo.RedirectStandardError = false;

                 process.StartInfo = startInfo;
                 process.Start();

                 process.WaitForExit();

             });
        }





        public async Task CheckAndDownloadKubectl()
        {

           

            if (!File.Exists(Path.Combine(FileSystemTools.getKonciergePath(), _kubectl))) {

                await DownloadKubeCtl();

            }
            //



        }

            public async Task DownloadKubeCtl()
        {


            string kubeVersionUrl = "https://dl.k8s.io/release/stable.txt"; // Replace with your URL
            try
            {
                using (HttpClient client = new HttpClient())
                {
                    // Fetch the content of the text file
                    string kubeVersion = await client.GetStringAsync(kubeVersionUrl);
                    Console.WriteLine("Latest kubectl version: " + kubeVersion);

                    

                   


                    var downloadUrl = $"https://dl.k8s.io/release/{kubeVersion.Trim()}/bin/{_osName}/amd64/{_kubectl}";

                    // Now use the content in another URL to download a file
                    Console.WriteLine("Downloading Kubectl: " + downloadUrl);

                    using (var clientD = new WebClient())
                    {
                        string destFile = Path.Combine(FileSystemTools.getKonciergePath(), _kubectl);
                        clientD.DownloadFile(downloadUrl, destFile);
                    }
                }


                if (_osName == "linux")
                {
                    Console.WriteLine("Making Kubectl an executable ");


                string command = $"chmod +x {Path.Combine(FileSystemTools.getKonciergePath(), _kubectl)}";

                System.Diagnostics.Process process = new System.Diagnostics.Process();
                System.Diagnostics.ProcessStartInfo startInfo = new System.Diagnostics.ProcessStartInfo();
                startInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Normal;
                startInfo.FileName = "/bin/bash";
                startInfo.WorkingDirectory = FileSystemTools.getKonciergePath();
                startInfo.Arguments = $"-c \"{command}\"";
                process.StartInfo = startInfo;
                    startInfo.RedirectStandardOutput = false;
                    startInfo.RedirectStandardInput = false;
                    startInfo.RedirectStandardError = true;
                    process.Start();
                   
                process.WaitForExit();
                }
                
                
            }
            catch (Exception ex)
            {
                Console.WriteLine("Error: " + ex.Message);
            }
            
            
        }

            
            

        public Guid ConnectTo(KubeConfigEntity kc, string? contextName)
        {
            var newConn = new KonciergeConnectionModel(kc, contextName);
            _connections.Add(newConn);
            return newConn.id;
        }

        public string buildForwardCommandString(ForwardEntity tmpFwd)
        {
            return $"port-forward --kubeconfig=\"{tmpFwd.WithConfig.Path}\" --context=\"{tmpFwd.Context}\" --namespace=\"{tmpFwd.Namespace}\" {(tmpFwd.Type==KonciergeForwardType.Service ? "service" : "pod")}/{tmpFwd.Selector} {tmpFwd.LocalPort}:{tmpFwd.RemotePort}";
        }

        public void SaveForward(ForwardEntity tmpFwd)
        {
            _kubeForwardRepository.Create(tmpFwd).Wait();
            
        }

        private KonciergeConnectionModel KonciergeConnection(Guid id) =>_connections.First(x=>x.id==id);
        

        public List<KubeNamespaceDto> GetNamespacesForConnection(Guid id)
        {

          



            return KonciergeConnection(id).GetKubeNamespaces();
        }

        


        public async Task CheckAndAddDefaultConfig() {

            if (_kubeConfigRepository.DefaultExists())
                return;

            var kubeConfigPath = KubernetesClientConfiguration.KubeConfigDefaultLocation;
            _logger.LogError($"Looking for Default KubeConfig @ {kubeConfigPath}");

            if (File.Exists(kubeConfigPath)) {

                var exists = _kubeConfigRepository.DefaultExists();

                var newConfig = new KubeConfigEntity() { 
                
                    Id = Guid.NewGuid(),
                    Name = Path.GetFileName(kubeConfigPath),
                    Path = kubeConfigPath,
                    IsDefault = true

                };

               await _kubeConfigRepository.Create(newConfig);


            }

            

        }

        #region KubeConfig
        public KubeConfigEntity GetKubeConfig(Guid? kubeConfigId)
        {

            if (kubeConfigId.HasValue)
            {

                var defaultKc = _kubeConfigRepository.getDefaultKubeconfig(true).Result;

                if (defaultKc is null)
                {

                    throw new KubeConfigNotFoundException("default");

                }
                return defaultKc;

            }
            else
            {

                var specificKc = _kubeConfigRepository.GetById(kubeConfigId!.Value, true).Result;


                if (specificKc is null)
                {

                    throw new KubeConfigNotFoundException(kubeConfigId.Value.ToString());

                }
                return specificKc;
            }


           

        }

        public IQueryable<KubeConfigEntity> GetKubeConfigs()=> _kubeConfigRepository.GetAll(false);

       
   


    #endregion

    #region Context

    public List<KubeContextDto> GetContextsForConfig(Guid id)
    {
            var ret=new List<KubeContextDto>();

            var kc= GetKubeConfig(id);

            

            var config= KubernetesYaml.LoadFromFileAsync<K8SConfiguration>(kc.Path).Result;

            var def = config.CurrentContext;

            foreach (var context in config.Contexts)
            {
                ret.Add(
                    new KubeContextDto() {
                    
                        Name = context.Name,
                        IsCurrent=def is not null && def== context.Name

                    }
                    );
            }

            return ret;

    }

      


        #endregion
        #region Services

        public List<KubeForwardableDto> GetServicesForNamespace(Guid idConn, string tmpFwdNamespace) =>
            KonciergeConnection(idConn).GetKubeServices(tmpFwdNamespace);

        #endregion
        
        #region Pods

        public List<KubeForwardableDto> GetPodsForNamespace(Guid idConn, string tmpFwdNamespace) =>
            KonciergeConnection(idConn).GetKubePods(tmpFwdNamespace);

       

        #endregion

        
        #region Configs

        public KubeConfigDto GetKubeMapAndSecretsForNamespace(Guid idConn, string tmpFwdNamespace) =>
            KonciergeConnection(idConn).GetKubeMapAndSecretsForNamespace(tmpFwdNamespace);

        #endregion


        public IQueryable<ForwardEntity> GetAllForwards() => _kubeForwardRepository.GetAllWithInclude();



    }
}