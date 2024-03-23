using System.Collections.Generic;

namespace FreightCMS.Models;

// see https://kb.freightcms.com/packaging/commodities/
public class CommodityModel
{
    public required string Id { get; set; }
    public required string Description { get; set; }
    public WeightModel? Weight { get; set; }
    public DimensionsModel? Dimensions { get; set; }
    public VolumeModel? Volume { get; set; }
    public string? NFMC { get; set; }
    public string? HarmonizedCode { get; set; }
    public string? CountryOfOrigin { get; set; }
    public CurrencyModel? DeclaredValue { get; set; }
    public bool IsHazardous { get; set; } = false;
    public bool IsStackable { get; set; } = false;
    public required Metadata Metadata { get; set; } = new Metadata();
    public TemperatureModel? MinTemperature { get; set; }
    public TemperatureModel? MaxTemperature { get; set; }
    public required ICollection<ReferenceModel> References { get; set; } = new List<ReferenceModel>();
}