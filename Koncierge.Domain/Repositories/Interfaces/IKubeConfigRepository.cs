using Koncierge.Domain.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace Koncierge.Domain.Repositories.Interfaces
{

    public interface IKubeConfigRepository : IGenericRepository<KubeConfigEntity>
    {
        bool DefaultExists();
        Task<KubeConfigEntity?> getDefaultKubeconfig(bool asReadonly= false);



    }
}
