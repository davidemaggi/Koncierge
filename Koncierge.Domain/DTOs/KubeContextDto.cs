using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Domain.DTOs
{
    public class KubeContextDto
    {

        public string Name { get; set; }
        public bool IsCurrent { get; set; }





        public override string ToString()
        {
            return this.Name;
        }
    }
}
