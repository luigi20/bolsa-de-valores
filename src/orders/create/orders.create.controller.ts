import { Body, Controller, Param, Post } from '@nestjs/common';
import { OrdersCreateService } from './orders.create.service';
import { InitTransactionDTO } from './order.dto';

@Controller('wallets/:wallet_id/orders')
export class OrdersCreateController {
  constructor(private readonly ordersCreateService: OrdersCreateService) {}

  @Post()
  public async create(
    @Param('wallet_id') wallet_id: string,
    @Body()
    body: InitTransactionDTO,
  ) {
    return await this.ordersCreateService.execute({
      ...body,
      wallet_id,
    });
  }
}
