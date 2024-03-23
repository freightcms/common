namespace FreightCMS.Models;

public class DimensionsModel
{
    public double Length { get; set; }
    public double Width { get; set; }
    public double Height { get; set; }
    public required string Unit { get; set; }
}