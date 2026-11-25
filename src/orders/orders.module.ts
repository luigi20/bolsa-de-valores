import { Module } from '@nestjs/common';
import { OrdersCreateService } from './create/orders.create.service';
import { OrdersCreateController } from './create/orders.create.controller';

@Module({
  controllers: [OrdersCreateController],
  providers: [OrdersCreateService],
})
export class OrdersModule {}
