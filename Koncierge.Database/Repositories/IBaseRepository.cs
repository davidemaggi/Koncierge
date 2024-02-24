using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Database.Repositories
{
    public interface IBaseRepository<T>
    {
        T AddOrUpdate(T data);
        T Create(T data);
        IEnumerable<T> All();
        T FindById(int id);
        T Update(T entity);
        bool Delete(int id);
        bool Exists(int id);
        bool Exists(T entity);
    }
}
