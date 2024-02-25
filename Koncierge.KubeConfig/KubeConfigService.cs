using k8s;
using k8s.KubeConfigModels;
using Koncierge.Exceptions;
using Koncierge.Models;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;
using static Koncierge.Models.Enums;

namespace Koncierge.KubeConfig
{
    public class KubeConfigService : IKubeConfigService
    {

        private string _defaultConfig = "config";
        public string kubePath;



        public KubeConfigService()
        {


            kubePath = Path.Join(Environment.GetFolderPath(Environment.SpecialFolder.UserProfile), ".kube");



        }

        public string GetKubeConfigDefaultPath() => kubePath;

        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath()
        {

            return await GetKubeConfigFromPath(kubePath);


        }
        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath(string p)
        {
            var ret = new List<KubeConfigFile>();

            foreach (string filePath in Directory.EnumerateFiles(p, "*.*", SearchOption.AllDirectories))
            {
                try
                {

                    var deserializer = new YamlDotNet.Serialization.DeserializerBuilder()
           .WithNamingConvention(CamelCaseNamingConvention.Instance)
           .IgnoreUnmatchedProperties()
           .Build();

                    K8SConfiguration kubeconfig = deserializer.Deserialize<K8SConfiguration>(File.ReadAllText(filePath));

                    var config = KubernetesClientConfiguration.BuildConfigFromConfigFile(filePath);

                    ret.Add(new KubeConfigFile(Path.GetFileName(filePath), filePath));

                }
                catch
                {


                }




            }


            return ret;
        }

        public async Task<KubeConfigFile> GetKubeConfigFileFromPath(string p)
        {
        

            if (!File.Exists(p)) { 
            throw new KubeConfigNotFountException();
            }

            if (!IsValidKubeConfig(p)) {
                throw KubeConfigNotValidException.WithKcf(p);
            }


            return new KubeConfigFile(Path.GetFileName(p), p);
        }

        public KubeConfigFileStatus CheckKubeConfig(KubeConfigFile toCheck)
        {


            if (!File.Exists(toCheck.Path))
            {
                return KubeConfigFileStatus.MISSING;
            }



            
            return IsValidKubeConfig(toCheck.Path) ? KubeConfigFileStatus.OK  : KubeConfigFileStatus.INVALID;
        }


        public bool IsValidKubeConfig(string toCheckPath) {

            try
            {

                var deserializer = new YamlDotNet.Serialization.DeserializerBuilder()
       .WithNamingConvention(CamelCaseNamingConvention.Instance)
       .IgnoreUnmatchedProperties()
       .Build();

                K8SConfiguration kubeconfig = deserializer.Deserialize<K8SConfiguration>(File.ReadAllText(toCheckPath));

                var config = KubernetesClientConfiguration.BuildConfigFromConfigFile(toCheckPath);


                return true;

            }
            catch
            {
                return false;

            }


        }

        public K8SConfiguration ReadKubeConfig(string filePath)
        {
           var deserializer = new YamlDotNet.Serialization.DeserializerBuilder()
           .WithNamingConvention(CamelCaseNamingConvention.Instance)
           .IgnoreUnmatchedProperties()
           .Build();

            K8SConfiguration kubeconfig = deserializer.Deserialize<K8SConfiguration>(File.ReadAllText(filePath));
            return kubeconfig;
        }

        public MergeResult MergeKubeConfig(string toBeMerged, string mergetTo, bool force = false, bool verbose = false)
        {


            var ret = new MergeResult();

            try
            {
                K8SConfiguration kubeconfig = ReadKubeConfig(mergetTo);
                K8SConfiguration kubeconfigMerge = ReadKubeConfig(toBeMerged);


                foreach (var ctx in kubeconfigMerge.Contexts)
                {

                    var contextAlreadyExists = kubeconfig.Contexts.Any(x => x.Name.Equals(ctx.Name, StringComparison.OrdinalIgnoreCase));


                    if (!contextAlreadyExists || force)
                    {
                        if (contextAlreadyExists)
                        {
                            ret.Modified.Add(ctx.Name);
                        }
                        else
                        {
                            ret.Added.Add(ctx.Name);
                        }

                        // Add Context
                        kubeconfig.Contexts = kubeconfig.Contexts.Where(x => !x.Name.Equals(ctx.Name, StringComparison.OrdinalIgnoreCase)).ToList();
                        kubeconfig.Contexts = kubeconfig.Contexts.Concat(new[] { ctx }).ToList();

                        //ret.details.Add(new MergeResultItemModel(Kind.context, ctx.Name, contextAlreadyExists ? EditAction.Modified : EditAction.Added));

                        var clusterAlreadyExists = kubeconfig.Clusters.Any(x => x.Name.Equals(ctx.ContextDetails.Cluster, StringComparison.OrdinalIgnoreCase));
                        if (clusterAlreadyExists)
                        {
                            kubeconfig.Clusters = kubeconfig.Clusters.Where(x => !x.Name.Equals(ctx.ContextDetails.Cluster, StringComparison.OrdinalIgnoreCase)).ToList();



                        }

                        // Add Cluster
                        var cluster = kubeconfigMerge.Clusters.FirstOrDefault(x => x.Name.Equals(ctx.ContextDetails.Cluster, StringComparison.OrdinalIgnoreCase));
                        kubeconfig.Clusters = kubeconfig.Clusters.Concat(new[] { cluster }).ToList();

                        // ret.details.Add(new MergeResultItemModel(Kind.cluster, ctx.ContextDetails.Cluster, clusterAlreadyExists ? EditAction.Modified : EditAction.Added));

                        // Add User

                        var userAlreadyExists = kubeconfig.Users.Any(x => x.Name.Equals(ctx.ContextDetails.User, StringComparison.OrdinalIgnoreCase));

                        if (userAlreadyExists)
                        {
                            kubeconfig.Users = kubeconfig.Users.Where(x => !x.Name.Equals(ctx.ContextDetails.User, StringComparison.OrdinalIgnoreCase)).ToList();

                        }

                        var user = kubeconfigMerge.Users.FirstOrDefault(x => x.Name.Equals(ctx.ContextDetails.User, StringComparison.OrdinalIgnoreCase));
                        kubeconfig.Users = kubeconfig.Users.Concat(new[] { user }).ToList();
                        //ret.details.Add(new MergeResultItemModel(Kind.user, ctx.ContextDetails.User, userAlreadyExists ? EditAction.Modified : EditAction.Added));






                    }



                }
                ret.Merged = kubeconfig;
            }
            catch (Exception e)
            {

                throw KubeConfigMergeFailedException.WithFromTo(toBeMerged, mergetTo);

            }

            return ret;

        }


        public Task SaveKubeConfig(string path, K8SConfiguration config, bool backup = true)
        {
            try
            {

                if (backup)
                {
                    File.Delete($"{path}.backup");
                    File.Copy(path, $"{path}.backup");
                }


                var serializer = new SerializerBuilder()
       .WithNamingConvention(CamelCaseNamingConvention.Instance)
       .Build();

                var stringResult = serializer.Serialize(config);

                File.WriteAllText(path, stringResult);




            }
            catch (Exception e)
            {

                throw new KubeConfigNotSavedException("Error saving config file");

            }

            return Task.CompletedTask;
        }

    }
}