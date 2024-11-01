using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Koncierge.Domain.DTOs
{
    public class KubeNamespaceDto
    {
        public string Name { get; set; }


        public override string ToString()
        {
            return Name;
        }

    }
}
