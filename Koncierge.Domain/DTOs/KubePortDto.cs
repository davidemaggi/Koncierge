namespace Koncierge.Domain.DTOs;

public class KubePortDto
{
     public int Port { get; set; }
     public string Protocol { get; set; }
     
     
     
     public override string ToString()
     {
          return $"{this.Port} ({Protocol})";
     }
     
     
}