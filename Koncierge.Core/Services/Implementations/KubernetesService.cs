using k8s;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using static Microsoft.EntityFrameworkCore.DbLoggerCategory.Database;

namespace Koncierge.Core.Services.Implementations
{
    public class KubernetesService : IKubernetesService
    {
        private IKubernetes _k8sClient;


        private KubernetesClientConfiguration _k8SClientConfig;


        public KubernetesService(KubeConfigEntity kc, string? contextname)
        {


            if (contextname is null)
            {

                _k8SClientConfig = KubernetesClientConfiguration.BuildConfigFromConfigFile(kc.Path);

            }
            else
            {

                _k8SClientConfig = KubernetesClientConfiguration.BuildConfigFromConfigFile(kc.Path, contextname);
                
            }

            _k8SClientConfig.SkipTlsVerify=true;
            _k8sClient = new Kubernetes(_k8SClientConfig);
            



        }

      

        public List<KubeNamespaceDto> GetNamespaces()
        {
            var ret = new List<KubeNamespaceDto>();
            var namespaces = _k8sClient.CoreV1.ListNamespace();

            foreach (var item in namespaces.Items)
            {
                ret.Add(new KubeNamespaceDto() { Name = item.Metadata.Name });
            }

            return ret;
        }

     
        public List<KubeForwardableDto> GetServices(string tmpFwdNamespace)
        {
            var ret = new List<KubeForwardableDto>();
            var services = _k8sClient.CoreV1.ListNamespacedService(tmpFwdNamespace);

            foreach (var item in services.Items)
            {
                var fwd = new KubeForwardableDto() { Name = item.Metadata.Name };
                
                fwd.Ports = new List<KubePortDto>();
                
                foreach (var port in item.Spec.Ports)
                {
                    fwd.Ports.Add(new KubePortDto(){Port = port.Port, Protocol = port.Protocol});
                }
                if(fwd.Ports.Count > 0)
                    ret.Add(fwd);
            }

            return ret;
        }
        
        
        public List<KubeForwardableDto> GetPods(string tmpFwdNamespace)
        {
            var ret = new List<KubeForwardableDto>();
            var services = _k8sClient.CoreV1.ListNamespacedPod(tmpFwdNamespace);

            foreach (var item in services.Items)
            {
                var fwd = new KubeForwardableDto() { Name = item.Metadata.Name };
                
                fwd.Ports = new List<KubePortDto>();
                
                foreach (var container in item.Spec.Containers)
                {
                    if (container.Ports is not null)
                    {
                    

                    foreach (var port in container.Ports)
                    {
                        fwd.Ports.Add(new KubePortDto(){Port = port.ContainerPort, Protocol = port.Protocol});
                    }}
                }
                if(fwd.Ports.Count > 0)
                ret.Add(fwd);
            }

            return ret;
        }

        public KubeConfigDto GetMapsAndSecrets(string tmpFwdNamespace)
        {
            var ret= new KubeConfigDto();

            var maps = _k8sClient.CoreV1.ListNamespacedConfigMap(tmpFwdNamespace);
            
            foreach (var map in maps)
            {
                ret.ConfigMaps.Add(new ConfigItemDto(map.Metadata.Name, map.Data.ToDictionary()));
            }
            
            

            var secrets = _k8sClient.CoreV1.ListNamespacedSecret(tmpFwdNamespace);

            
            foreach (var secret in secrets)
            {
                var newDict = new Dictionary<string, string>();
                
                foreach (var keyValuePair in secret.Data)
                {
                    
                    string base64EncodedString = Convert.ToBase64String(keyValuePair.Value);

// Decode the Base64 string to get the original bytes
                    byte[] decodedBytes = Convert.FromBase64String(base64EncodedString);

// Convert the decoded bytes into the actual string
                    string decodedString = System.Text.Encoding.UTF8.GetString(decodedBytes);
                    
                    newDict[keyValuePair.Key] = decodedString;
                }
                
                
                ret.Secrets.Add(new ConfigItemDto(secret.Metadata.Name, newDict));
            }

            
            return ret;
        }
    }

    
}
