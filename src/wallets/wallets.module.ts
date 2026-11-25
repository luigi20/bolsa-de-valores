import { Module } from '@nestjs/common';
import { WalletsCreateController } from './create/wallets.create.controller';
import { WalletsCreateService } from './create/wallets.create.service';
import { WalletsGetAllController } from './get_all/wallets.get_all.controller';
import { WalletsGetAllService } from './get_all/wallets.get_all.service';

@Module({
  controllers: [WalletsCreateController, WalletsGetAllController],
  providers: [WalletsCreateService, WalletsGetAllService],
})
export class WalletsModule {}
