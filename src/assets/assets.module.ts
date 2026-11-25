import { Module } from '@nestjs/common';
import { AssetsCreateController } from './create/assets.create.controller';
import { AssetsCreateService } from './create/assets.create.service';
import { AssetsGetAllController } from './get_all/assets.get_all.controller';
import { AssetsGetAllService } from './get_all/assets.get_all.service';

@Module({
  controllers: [AssetsCreateController, AssetsGetAllController],
  providers: [AssetsCreateService, AssetsGetAllService],
})
export class AssetsModule {}
