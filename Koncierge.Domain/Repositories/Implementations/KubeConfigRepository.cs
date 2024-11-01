using Koncierge.Domain.Entities;
using Koncierge.Domain.Repositories.Interfaces;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Domain.Repositories.Implementations
{
    public class KubeConfigRepository : GenericRepository<KubeConfigEntity>, IKubeConfigRepository
    {

        private readonly KonciergeContext _ctx;
        public KubeConfigRepository(KonciergeContext ctx) : base(ctx)

        {
            _ctx = ctx;
        }

        public  bool DefaultExists() => _ctx.KubeConfigs.Any(x=>x.IsDefault);
        

        public Task<KubeConfigEntity?> getDefaultKubeconfig(bool asReadonly = false) => asReadonly ? _ctx.KubeConfigs.AsNoTracking().Where(x=>x.IsDefault).FirstOrDefaultAsync() : _ctx.KubeConfigs.Where(x => x.IsDefault).FirstOrDefaultAsync();


    }
}
