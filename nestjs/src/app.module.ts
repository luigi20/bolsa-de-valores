import { Module } from '@nestjs/common';
import { PrismaModule } from './infra/prisma.module';
import { AssetsModule } from './assets/assets.module';
import { WalletsModule } from './wallets/wallets.module';
import { WalletAssetsModule } from './wallet_assets/wallet_assets.module';
import { OrdersModule } from './orders/orders.module';

@Module({
  imports: [
    OrdersModule,
    PrismaModule,
    AssetsModule,
    WalletsModule,
    WalletAssetsModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
