using Koncierge.Models;
using LiteDB;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Database.Repositories
{
    public abstract class BaseRepository<T> : IBaseRepository<T> where T : IBaseEntity
    {
        public ILiteDatabase DB { get; }
        public ILiteCollection<T> Collection { get; }

        protected BaseRepository(ILiteDatabase db)
        {
            DB = db;
            Collection = db.GetCollection<T>();
        }

        public virtual T AddOrUpdate(T entity)
        {
            if (Exists(entity))
            {

                return Create(entity);

            }
            else { 
                return Update(entity);

            }
        }

        public virtual T Create(T entity)
        {
            var now = DateTime.Now;
            entity.Added = now;
            entity.Modified = now;

            Collection.EnsureIndex(x => x.Id);

            var newId = Collection.Insert(entity);
            return Collection.FindById(newId.AsInt32);
        }

        public virtual IEnumerable<T> All()
        {
            return Collection.FindAll();
        }

        public virtual T FindById(int id)
        {
            return Collection.FindById(id);
        }

        public virtual T Update(T entity)
        {
            entity.Modified=DateTime.Now;
            Collection.Upsert(entity);
            return entity;
        }

        public virtual bool Delete(int id)
        {
            return Collection.Delete(id);
        }

        public virtual bool Exists(int id)
        {
            return Collection.FindById(id) is not null;
        }
        public virtual bool Exists(T entity)
        {
            return Collection.FindById(entity.Id) is not null;
        }
    }
}
