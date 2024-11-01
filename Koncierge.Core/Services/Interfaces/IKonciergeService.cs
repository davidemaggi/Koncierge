using Koncierge.Domain.DTOs;
using Koncierge.Domain.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Core.Services.Interfaces
{
    public interface IKonciergeService
    {

        public Guid ConnectTo(KubeConfigEntity kc, string? contextName);

        #region KubeConfig

        public IQueryable<KubeConfigEntity> GetKubeConfigs();

        #endregion

        #region Contexts

        public List<KubeContextDto> GetContextsForConfig(Guid id);
        public List<KubeNamespaceDto> GetNamespacesForConnection(Guid id);

        #endregion

        #region Services

        public List<KubeForwardableDto> GetServicesForNamespace(Guid idConn, string tmpFwdNamespace);


            #endregion
            #region Pods

            public List<KubeForwardableDto> GetPodsForNamespace(Guid idConn, string tmpFwdNamespace);


            #endregion

            public string buildForwardCommandString(ForwardEntity tmpFwd);
        public Task ExecuteKubectl(string command);
        public void SaveForward(ForwardEntity tmpFwd);
        public KubeConfigDto GetKubeMapAndSecretsForNamespace(Guid idConn, string tmpFwdNamespace);
        public IQueryable<ForwardEntity> GetAllForwards();
    }
}
