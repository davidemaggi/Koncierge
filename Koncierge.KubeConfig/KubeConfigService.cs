using k8s;
using k8s.KubeConfigModels;
using Koncierge.Exceptions;
using Koncierge.Models.Config;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using YamlDotNet.Serialization.NamingConventions;

namespace Koncierge.KubeConfig
{
    public class KubeConfigService: IKubeConfigService
    {

        private string _defaultConfig ="config";
        public string kubePath;



        public KubeConfigService() {


            kubePath = Path.Join(Environment.GetFolderPath(Environment.SpecialFolder.UserProfile), ".kube");

           

        }

        public string GetKubeConfigDefaultPath() => kubePath;

        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath() {

            return await GetKubeConfigFromPath(kubePath);


        }
        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath(string p)
        {
            var ret = new List<KubeConfigFile>();

            foreach (string filePath in Directory.EnumerateFiles(p, "*.*", SearchOption.AllDirectories))
            {
                try {

                    var deserializer = new YamlDotNet.Serialization.DeserializerBuilder()
           .WithNamingConvention(CamelCaseNamingConvention.Instance)
           .IgnoreUnmatchedProperties()
           .Build();

                    K8SConfiguration kubeconfig = deserializer.Deserialize<K8SConfiguration>(File.ReadAllText(filePath));

                    var config = KubernetesClientConfiguration.BuildConfigFromConfigFile(filePath);

                    ret.Add(new KubeConfigFile(Path.GetFileName(filePath),filePath));

                }
                catch {


                }




            }


            return ret;
        }


    }
}
