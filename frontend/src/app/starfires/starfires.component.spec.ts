import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StarfiresComponent } from './starfires.component';

describe('StarfiresComponent', () => {
  let component: StarfiresComponent;
  let fixture: ComponentFixture<StarfiresComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ StarfiresComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(StarfiresComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
