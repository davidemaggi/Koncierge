namespace Koncierge.Domain.DTOs;

public class KubeConfigDto
{
    public List<ConfigItemDto> ConfigMaps { get; set; }
    public List<ConfigItemDto> Secrets { get; set; }


    public KubeConfigDto()
    {
        
        ConfigMaps = new List<ConfigItemDto>();
        Secrets = new List<ConfigItemDto>();
        
    }


}

public class ConfigItemDto
{
    public string Selector { get; set; }
    public Dictionary<string, string> Values { get; set; }

    public ConfigItemDto(string name, Dictionary<string, string> dictionary)
    {
        Selector = name;
        Values = dictionary;
    }
    public override string ToString()
    {
        return this.Selector;
    }
}



