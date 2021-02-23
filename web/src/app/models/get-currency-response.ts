import {CurrencyModel} from './currency-model';

export interface GetCurrencyResponse {
  status: boolean;
  message: string;
  data: CurrencyModel;
}
