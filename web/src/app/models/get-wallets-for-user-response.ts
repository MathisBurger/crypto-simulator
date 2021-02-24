import {CurrencyArrayEntry} from './currency-array-entry';

export interface GetWalletsForUserResponse {
  status: boolean;
  message: string;
  data: CurrencyArrayEntry[];
}
