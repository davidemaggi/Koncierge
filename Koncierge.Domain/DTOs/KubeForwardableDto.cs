namespace Koncierge.Domain.DTOs;

public class KubeForwardableDto
{
    public string Name { get; set; }
    public  List<KubePortDto> Ports { get; set; }





    public override string ToString()
    {
        return this.Name;
    }
}