using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Core.Services.Interfaces
{
    public interface IKubernetesService
    {
        public List<KubeNamespaceDto> GetNamespaces();
        public List<KubeForwardableDto> GetServices(string tmpFwdNamespace);
        public List<KubeForwardableDto> GetPods(string tmpFwdNamespace);


        public KubeConfigDto GetMapsAndSecrets(string tmpFwdNamespace);
    }
}
