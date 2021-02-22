import {CurrencyModel} from './currency-model';

export interface GetAllCurrencysResponse {
  status: boolean;
  message: string;
  data: CurrencyModel[];
}
