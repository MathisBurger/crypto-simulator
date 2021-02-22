export interface CurrencyModel {
  id: number;
  coin_id: string;
  rank: number;
  symbol: string;
  name: string;
  supply: number;
  max_supply: number;
  market_cap_usd: number;
  volume_usd_24_hr: number;
  price_usd: number;
  change_percent_24_hr: number;
  vwap_24_hr: number;
}
