using System.Collections.Generic;

namespace FreightCMS.Models;

public class HandlingUnitModel
{
    public required string Id { get; set; }
    public required string Description { get; set; }
    public WeightModel Weight { get; set; } = new WeightModel();
    public DimensionsModel Dimensions { get; set; } = new DimensionsModel();
    public VolumeModel Volume { get; set; } = new VolumeModel();
    public string? NFMC { get; set; }
    public string? HarmonizedCode { get; set; }
    public string? CountryOfOrigin { get; set; }
    public CurrencyModel? DeclaredValue { get; set; }
    public bool IsHazardous { get; set; } = false;
    public bool IsStackable { get; set; } = false;
    public Metadata Metadata { get; set; } = new Metadata();
    public TemperatureModel? MinTemperature { get; set; } = new TemperatureModel();
    public TemperatureModel? MaxTemperature { get; set; } = new TemperatureModel();
    public ICollection<ReferenceModel> References { get; set; } = new List<ReferenceModel>();
    public ICollection<CommodityModel> Commodities { get; set; } = new List<CommodityModel>();
}