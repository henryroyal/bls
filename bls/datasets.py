import io
import csv
import requests


# https://www.bls.gov/help/hlpforma.htm#CE
class DataSet:
    def __init__(self, symbol, name):
        self.name = name
        self.symbol = symbol
        self.url = 'https://download.bls.gov/pub/time.series/' + symbol

    @property
    def documentation_url(self):
        return self.url + '/{}.txt'.format(self.symbol)

    @property
    def series_index_url(self):
        return self.url + '/{}.series'.format(self.symbol)

    def __repr__(self):
        return '"{}": "{}"'.format(self.name, self.documentation_url)

    def get_series_index(self):
        response = requests.get(self.series_index_url)  # TODO: handle errors
        return csv.DictReader(io.StringIO(response.text), delimiter='\t')


#
#
# Employment & Unemployment
#
#

class NationalEmploymentHoursAndEarnings(DataSet):
    def __init__(self):
        super().__init__('ce', 'National Employment, Hours, and Earnings')


class NationalEmploymentHoursAndEarningsSIC(DataSet):
    def __init__(self):
        super().__init__('ee', 'National Employment, Hours, and Earnings (SIC basis)')


class StateAndAreaEmploymentHoursAndEarnings(DataSet):
    def __init__(self):
        super().__init__('sm', 'State and Area Employment, Hours, and Earnings')


class StateAndAreaEmploymentHoursAndEarningsSIC(DataSet):
    def __init__(self):
        super().__init__('sa', 'State and Area Employment, Hours, and Earnings (SIC basis)')


class StateAndCountyEmploymentAndWagesQuarterly(DataSet):
    def __init__(self):
        super().__init__('en', 'State and County Employment and Wages from Quarterly Census of Employment and Wages')


class OccupationalEmploymentStatistics(DataSet):
    def __init__(self):
        super().__init__('oe', 'Occupational Employment Statistics (OES)')


class BusinessEmploymentDynamics(DataSet):
    def __init__(self):
        super().__init__('bd', 'Business Employment Dynamics')


class LocalAreaUnemploymentStatistics(DataSet):
    def __init__(self):
        super().__init__('la', 'Local Area Unemployment Statistics')


class GeographicProfile(DataSet):
    def __init__(self):
        super().__init__('gp', 'Geographic Profile')


class MassLayoffStatistics(DataSet):
    def __init__(self):
        super().__init__('ml', 'Mass Layoff Statistics')


class JobOpeningsAndLaborTurnoverSIC(DataSet):
    def __init__(self):
        super().__init__('jl', 'Job Openings and Labor Turnover Survey (SIC basis)')


class JobOpeningsAndLaborTurnover(DataSet):
    def __init__(self):
        super().__init__('jt', 'Job Openings and Labor Turnover Survey')


class GreenGoodsAndServices(DataSet):
    def __init__(self):
        super().__init__('gg', 'Green Goods and Services')


#
#
# Spending & Time Use
#
#

class AmericanTimeUseSurvey(DataSet):
    def __init__(self):
        super().__init__('tu', 'American Time Use Survey')


class ConsumerExpenditureSurvey(DataSet):
    def __init__(self):
        super().__init__('cx', 'Consumer Expenditure Survey')


#
#
# Inflation & Prices
#
#

class AveragePriceData(DataSet):
    def __init__(self):
        super().__init__('ap', 'Average Prices')


class ConsumerPriceIndexAllUrban(DataSet):
    def __init__(self):
        super().__init__('cu', 'Consumer Price Index - All Urban Consumers')


class ConsumerPriceIndexUrbanWageWorkers(DataSet):
    def __init__(self):
        super().__init__('cw', 'Consumer Price Index - Urban Wage Earners and Clerical Workers')


class LegacyConsumerPriceIndexAllUrban(DataSet):
    def __init__(self):
        super().__init__('mu', 'Consumer Price Index - All Urban Consumers (Old Series)')


class LegacyConsumerPriceIndexUrbanWageWorkers(DataSet):
    def __init__(self):
        super().__init__('mw', 'Consumer Price Index - Urban Wage Earners and Clerical Workers (Old Series)')


class DepartmentStoreInventoryPriceIndex(DataSet):
    def __init__(self):
        super().__init__('li', 'Department Store Inventory Price Index')


class ChainedConsumerPriceIndexAllUrban(DataSet):
    def __init__(self):
        super().__init__('su', 'Chained CPI-All Urban Consumers')


class ProducerPriceIndexIndustryData(DataSet):
    def __init__(self):
        super().__init__('pc', 'Producer Price Index Industry Data - Current Series')


class LegacyProducerPriceIndexIndustryDataNAICS(DataSet):
    def __init__(self):
        super().__init__('nd', 'Producer Price Index Industry Data - Discontinued Series (NAICS basis)')


class LegacyProducerPriceIndexIndustryDataSIC(DataSet):
    def __init__(self):
        super().__init__('pd', 'Producer Price Index Industry Data - Discontinued Series (SIC basis)')


class ProducerPriceIndexCommodityData(DataSet):
    def __init__(self):
        super().__init__('wp', 'Producer Price Index Commodity Data - Current Series')


class LegacyProducerPriceIndexCommodityData(DataSet):
    def __init__(self):
        super().__init__('wd', 'Producer Price Index Commodity Data - Discontinued Series')


#
#
# Pay & Benefits
#
#

class WorkStoppageData(DataSet):
    def __init__(self):
        super().__init__('ws', 'Work Stoppage Data')


class EmployeeBenefitsSurvey(DataSet):
    def __init__(self):
        super().__init__('eb', 'Employee Benefits Survey')


class EmploymentCostIndexSIC(DataSet):
    def __init__(self):
        super().__init__('ec', 'Employment Cost Index (SIC)')


class EmployerCostForEmployeeCompensationSIC(DataSet):
    # FIXME - there is a formatting error in the headers of https://download.bls.gov/pub/time.series/cc/cc.series
    def __init__(self):
        super().__init__('cc', 'Employer Costs for Employee Compensation')


class EmployerCostForEmployeeCompensationQuarterly(DataSet):
    def __init__(self):
        super().__init__('ci', 'Employer Costs for Employee Compensation by Quarter')


class NationalCompensationSurvey(DataSet):
    def __init__(self):
        super().__init__('nc', 'National Compensation Survey')


class BenefitsCompensationSurvey(DataSet):
    def __init__(self):
        super().__init__('nb', 'National Compensation Survey - Benefits (Beginning with 2010 data)')


class ModeledWageEstimates(DataSet):
    def __init__(self):
        super().__init__('wm', 'Modeled Wage Estimates')


#
#
# Productivity
#
#

class MajorSectorProductivityAndCosts(DataSet):
    def __init__(self):
        super().__init__('pr', 'Major Sector Productivity and Costs')


class MajorSectorMultifactorProductivity(DataSet):
    def __init__(self):
        super().__init__('mp', 'Major Sector Multifactor Productivity')


class IndustryProductivity(DataSet):
    def __init__(self):
        super().__init__('ip', 'Industry Productivity')


#
#
# Workplace Injuries
#
#

class OccupationalInjuriesAndIllnessesPost2014(DataSet):
    def __init__(self):
        super().__init__('is', 'Occupational Injuries and Illnesses - industry data (2014 forward)')


class FatalOccupationalInjuriesPost2011(DataSet):
    def __init__(self):
        super().__init__('fw', 'Census of Fatal Occupational Injuries (2011 forward)')


class FatalOccupationalInjuries2003to2010(DataSet):
    def __init__(self):
        super().__init__('fi', 'Census of Fatal Occupational Injuries (2003 - 2010)')


class FatalOccupationalInjuries1992to2002(DataSet):
    def __init__(self):
        super().__init__('cf', 'Census of Fatal Occupational Injuries (1992 - 2002)')


class NonfatalOccupationalInjuriesPost2011(DataSet):
    def __init__(self):
        super().__init__('cs', 'Nonfatal cases involving days away from work: selected characteristics (2011 forward)')


class NonfatalOccupationalInjuries2003to2010(DataSet):
    def __init__(self):
        super().__init__('ch', 'Nonfatal cases involving days away from work: selected characteristics (2003 - 2010)')


class NonfatalOccupationalInjuries2002(DataSet):
    def __init__(self):
        super().__init__('hc', 'Nonfatal cases involving days away from work: selected characteristics (2002)')


class NonfatalOccupationalInjuries1992to2001(DataSet):
    def __init__(self):
        super().__init__('cd', 'Nonfatal cases involving days away from work: selected characteristics (1992 - 2001)')


class NonfatalOccupationalInjuriesAndIllnessesPre1989(DataSet):
    def __init__(self):
        super().__init__('hs', 'Occupational injuries and illnesses: industry data (pre-1989)')


class NonfatalOccupationalInjuriesAndIllnesses1989to2001(DataSet):
    def __init__(self):
        super().__init__('sh', 'Occupational injuries and illnesses: industry data (1989 - 2001)')


class NonfatalOccupationalInjuriesAndIllnesses2002(DataSet):
    def __init__(self):
        super().__init__('si', 'Occupational injuries and illnesses: industry data (2002)')


class NonfatalOccupationalInjuriesAndIllnesses2003to2013(DataSet):
    def __init__(self):
        super().__init__('ii', 'Occupational Injuries and Illnesses - industry data (2003 - 2013)')


#
#
# Occupational Requirements
#
#

class OccupationalRequirementsSurvey(DataSet):
    def __init__(self):
        super().__init__('or', 'Occupational Requirements Survey')


#
#
# International
#
#

class ImportExportPriceIndexes(DataSet):
    def __init__(self):
        super().__init__('ei', 'Import/Export Price Indexes')


timeseries = [

    # Employment & Unemployment
    NationalEmploymentHoursAndEarnings(),
    NationalEmploymentHoursAndEarningsSIC(),
    StateAndAreaEmploymentHoursAndEarnings(),
    StateAndAreaEmploymentHoursAndEarningsSIC(),
    StateAndCountyEmploymentAndWagesQuarterly(),
    OccupationalEmploymentStatistics(),
    BusinessEmploymentDynamics(),
    LocalAreaUnemploymentStatistics(),
    GeographicProfile(),
    MassLayoffStatistics(),
    JobOpeningsAndLaborTurnoverSIC(),
    JobOpeningsAndLaborTurnover(),
    GreenGoodsAndServices(),

    # Spending & Time Use
    AmericanTimeUseSurvey(),
    ConsumerExpenditureSurvey(),

    # Inflation & Prices
    AveragePriceData(),
    ConsumerPriceIndexAllUrban(),
    ConsumerPriceIndexUrbanWageWorkers(),
    LegacyConsumerPriceIndexAllUrban(),
    LegacyConsumerPriceIndexUrbanWageWorkers(),
    DepartmentStoreInventoryPriceIndex(),
    ChainedConsumerPriceIndexAllUrban(),
    ProducerPriceIndexIndustryData(),
    LegacyProducerPriceIndexIndustryDataNAICS(),
    LegacyProducerPriceIndexIndustryDataSIC(),
    ProducerPriceIndexCommodityData(),
    LegacyProducerPriceIndexCommodityData(),

    # Pay & Benefits
    WorkStoppageData(),
    EmployeeBenefitsSurvey(),
    EmploymentCostIndexSIC(),
    # EmployerCostForEmployeeCompensationSIC(),  # FIXME
    EmployerCostForEmployeeCompensationQuarterly(),
    NationalCompensationSurvey(),
    BenefitsCompensationSurvey(),
    ModeledWageEstimates(),

    # Productivity
    MajorSectorProductivityAndCosts(),
    MajorSectorMultifactorProductivity(),
    IndustryProductivity(),

    # Workplace Injuries
    OccupationalInjuriesAndIllnessesPost2014(),
    FatalOccupationalInjuriesPost2011(),
    FatalOccupationalInjuries2003to2010(),
    FatalOccupationalInjuries1992to2002(),
    NonfatalOccupationalInjuriesPost2011(),
    NonfatalOccupationalInjuries2003to2010(),
    NonfatalOccupationalInjuries2002(),
    NonfatalOccupationalInjuries1992to2001(),
    NonfatalOccupationalInjuriesAndIllnessesPre1989(),
    NonfatalOccupationalInjuriesAndIllnesses1989to2001(),
    NonfatalOccupationalInjuriesAndIllnesses2002(),
    NonfatalOccupationalInjuriesAndIllnesses2003to2013(),

    # Occupational Requirements
    OccupationalRequirementsSurvey(),

    # International
    ImportExportPriceIndexes(),
]
