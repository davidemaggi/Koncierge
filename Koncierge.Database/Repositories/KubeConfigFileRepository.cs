using Koncierge.Models;
using LiteDB;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Numerics;
using System.Text;
using System.Threading.Tasks;
using static System.Runtime.InteropServices.JavaScript.JSType;

namespace Koncierge.Database.Repositories
{


    public class KubeConfigFileRepository : BaseRepository<KubeConfigFile>, IKubeConfigFileRepository
    {
        public KubeConfigFileRepository(ILiteDatabase db)
        : base(db)
        { }


        public override KubeConfigFile AddOrUpdate(KubeConfigFile entity)
        {
            if (!Exists(entity.Path))
            {

                return Create(entity);

            }
            else
            {
                var tmp = FindByPath(entity.Path);
                entity.Id = tmp.Id;
                entity.JustAdded = false;
                return Update(entity);

            }
        }


        public bool Exists(string path)
        {


            
            return All().Any(x => x.Path == path);


        }
        public KubeConfigFile FindByPath(string path)
        {

            var ret = All().Where(x => x.Path == path).First();

            return ret;


        }
        public void ClearJustAdded()
        {

            foreach (var config in All()) {


                config.JustAdded = false;
                AddOrUpdate(config);

            }

            


        }

    }

}
