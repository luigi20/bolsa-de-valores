import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { OrdersGetService } from './orders.get.service';

@Controller('wallets/:wallet_id/orders')
export class OrdersGetController {
  constructor(private readonly ordersGetService: OrdersGetService) {}

  @Get()
  public async get(@Param('wallet_id') wallet_id: string) {
    return this.ordersGetService.execute(wallet_id);
  }
}
