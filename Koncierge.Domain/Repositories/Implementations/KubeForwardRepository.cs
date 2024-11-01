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
    public class KubeCForwardRepository : GenericRepository<ForwardEntity>, IKubeForwardRepository
    {

        private readonly KonciergeContext _ctx;
        public KubeCForwardRepository(KonciergeContext ctx) : base(ctx)

        {
            _ctx = ctx;
        }

        public IQueryable<ForwardEntity> GetAllWithInclude()
        {
            var ret = _ctx.Set<ForwardEntity>().Include(x=>x.WithConfig).Include(x=>x.Configs);
            return ret;
        }

        public IQueryable<ForwardEntity> GetAllWithIncludeForConfig(Guid confId, string context)
        {
            return _ctx.Set<ForwardEntity>().Include(x=>x.WithConfig).Where(x=>x.WithConfig.Id==confId && x.Context==context);
        }
    }
}
