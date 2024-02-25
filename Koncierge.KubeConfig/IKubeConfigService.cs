using k8s.KubeConfigModels;
using Koncierge.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using static Koncierge.Models.Enums;

namespace Koncierge.KubeConfig
{
    public interface  IKubeConfigService
    {

        public string GetKubeConfigDefaultPath();
        public Task<List<KubeConfigFile>> GetKubeConfigFromPath();
        public Task<List<KubeConfigFile>> GetKubeConfigFromPath(string p);
        public KubeConfigFileStatus CheckKubeConfig(KubeConfigFile toCheck);

        public bool IsValidKubeConfig(string toCheckPath);
        public Task<KubeConfigFile> GetKubeConfigFileFromPath(string p);
        public MergeResult MergeKubeConfig(string toBeMerged, string mergetTo, bool force = false, bool verbose = false);
        public K8SConfiguration ReadKubeConfig(string filePath);
        public Task SaveKubeConfig(string path, K8SConfiguration config, bool backup = true);
    }
}
