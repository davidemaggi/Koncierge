using Koncierge.Core.Exceptions;
using Koncierge.Core.Services.Implementations;
using Koncierge.Core.Services.Interfaces;
using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using Koncierge.Domain.Enums;
using Koncierge.Domain.Repositories.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography.X509Certificates;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Core.Models
{
    public class KonciergeConnectionModel
    {
        public readonly Guid id;
        private readonly KubeConfigEntity _kubeConfig;
        private readonly IKubernetesService _kubernetesConnection;
        private KonciergeConnectionStatus Status;


        private readonly IKubeConfigRepository _kubeConfigRepository;
        public KonciergeConnectionModel(KubeConfigEntity kc, string? contextname)
        {
            id = Guid.NewGuid();
            _kubeConfig = kc;
            _kubernetesConnection=new KubernetesService(kc, contextname);

        }


        public List<KubeNamespaceDto> GetKubeNamespaces() => _kubernetesConnection.GetNamespaces();


        public List<KubeForwardableDto> GetKubeServices(string tmpFwdNamespace) =>
            _kubernetesConnection.GetServices(tmpFwdNamespace);

        public List<KubeForwardableDto> GetKubePods(string tmpFwdNamespace)=>
            _kubernetesConnection.GetPods(tmpFwdNamespace);


        public KubeConfigDto GetKubeMapAndSecretsForNamespace(string tmpFwdNamespace)=> _kubernetesConnection.GetMapsAndSecrets(tmpFwdNamespace);
        
    }
}
