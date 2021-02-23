import {TradeModel} from './trade-model';

export interface GetAllTradesResponse {
  status: boolean;
  message: string;
  data: TradeModel[];
}
