import { Test, TestingModule } from '@nestjs/testing';
import { AssetsGetAllController } from './assets.get_all.controller';
import { AssetsGetAllService } from './assets.get_all.service';

describe('AssetsController', () => {
  let controller: AssetsGetAllController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [AssetsGetAllController],
      providers: [AssetsGetAllService],
    }).compile();

    controller = module.get<AssetsGetAllController>(AssetsGetAllController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
