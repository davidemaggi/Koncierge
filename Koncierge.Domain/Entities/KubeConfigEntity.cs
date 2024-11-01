using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Domain.Entities
{
    public class KubeConfigEntity: BaseEntity
    {
        public string Name { get; set; }
        public string Path { get; set; }
        public bool IsDefault { get; set; }

        public ICollection<ForwardEntity> HasForwards { get; set; }




        public override string ToString()
        {
            return this.Name;
        }

    }
}
