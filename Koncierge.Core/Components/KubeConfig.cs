using Koncierge.Models;

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using YamlDotNet.Serialization.NamingConventions;
using YamlDotNet.Serialization;
using Koncierge.Exceptions;
using k8s.KubeConfigModels;

namespace Koncierge.Core
{
    public partial class KonciergeCoreService
    {

        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath(bool dry = false) { 
            
            
            var configs= await _kubeConfig.GetKubeConfigFromPath();

            if (!dry) { return await SaveConfigs(configs); }

            return configs;
        }


        public async Task<List<KubeConfigFile>> GetKubeConfigFromPath(string path, bool dry = false)  {


            var configs = await _kubeConfig.GetKubeConfigFromPath(path); 
        
            if (!dry) { SaveConfigs(configs); }

            return configs;
        }

        public string GetKubeConfigDefaultPath() => _kubeConfig.GetKubeConfigDefaultPath();

        private async Task<List<KubeConfigFile>> SaveConfigs(List<KubeConfigFile> configs) {

            foreach (var config in configs)
            {

                _konciergeDbService.KubeConfigFileRepository().AddOrUpdate(config);


            }
            var ret=_konciergeDbService.KubeConfigFileRepository().All().ToList();

            _konciergeDbService.KubeConfigFileRepository().ClearJustAdded();

            return ret;
        }


        public async Task<List<KubeConfigFile>> GetKubeConfigList()
        {
            return _konciergeDbService.KubeConfigFileRepository().All().ToList();
        }


        public async Task<bool> RemoveKubeConfig(int id)
        {
            return _konciergeDbService.KubeConfigFileRepository().Delete(id);
        }

        public async Task<bool> IsValidKubeConfig(string path)
        {
            return   _kubeConfig.IsValidKubeConfig(path);
        }

        public async Task<MergeResult> MergeKubeConfig(string from, string to)
        {
            return _kubeConfig.MergeKubeConfig(from,to);
        }
        public async Task<KubeConfigFile> GetKubeConfigFileFromPath(string path)
        {
            return await _kubeConfig.GetKubeConfigFileFromPath(path);
        }

        public Task SaveKubeConfig(string path, K8SConfiguration config, bool backup = true)
        {
            return _kubeConfig.SaveKubeConfig(path, config, backup);

        }

    }
}
