import { Body, Controller, Param, Post } from '@nestjs/common';
import { OrdersExecuteService } from './orders.execute.service';
import { OrderExecuteDTO } from './order_execute.dto';

@Controller('orders')
export class OrdersExecuteController {
  constructor(private readonly ordersExecuteService: OrdersExecuteService) {}

  @Post('execute')
  public async executeTransaction(@Body() body: OrderExecuteDTO) {
    return await this.ordersExecuteService.execute(body);
  }
}
