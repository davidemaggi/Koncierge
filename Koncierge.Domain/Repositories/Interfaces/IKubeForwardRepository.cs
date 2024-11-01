using Koncierge.Domain.Entities;

namespace Koncierge.Domain.Repositories.Interfaces;

public interface IKubeForwardRepository : IGenericRepository<ForwardEntity>
{
    
    IQueryable<ForwardEntity> GetAllWithInclude();
    IQueryable<ForwardEntity> GetAllWithIncludeForConfig(Guid confId, string context);
}