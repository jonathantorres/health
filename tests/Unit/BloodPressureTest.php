<?php

namespace Tests\Unit;

use App\BloodPressure;
use App\User;
use Carbon\Carbon;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Tests\TestCase;

class BloodPressureTest extends TestCase
{
    use RefreshDatabase;

    private User $user;

    public function setUp(): void
    {
        parent::setUp();

        $this->user = User::factory()->create();
    }

    /** @test */
    public function normal_blood_pressure_severity_is_calculated(): void
    {
        $normalReading = BloodPressure::create([
            'systolic' => 120,
            'diastolic' => 80,
            'pulse' => 0,
            'user_id' => $this->user->id,
            'reading_date' => Carbon::now(),
        ]);
        $this->assertSame($normalReading->severity()['text'], 'Normal');
    }

    /** @test */
    public function pre_blood_pressure_severity_is_calculated(): void
    {
        $preReading = BloodPressure::create([
            'systolic' => 130,
            'diastolic' => 85,
            'pulse' => 0,
            'user_id' => $this->user->id,
            'reading_date' => Carbon::now(),
        ]);
        $this->assertSame($preReading->severity()['text'], 'Pre Hypertension');
    }

    /** @test */
    public function stage1_blood_pressure_severity_is_calculated(): void
    {
        $stage1Reading = BloodPressure::create([
            'systolic' => 140,
            'diastolic' => 99,
            'pulse' => 0,
            'user_id' => $this->user->id,
            'reading_date' => Carbon::now(),
        ]);
        $this->assertSame($stage1Reading->severity()['text'], 'Stage 1 Hypertension');
    }

    /** @test */
    public function stage2_blood_pressure_severity_is_calculated(): void
    {
        $stage2Reading = BloodPressure::create([
            'systolic' => 161,
            'diastolic' => 101,
            'pulse' => 0,
            'user_id' => $this->user->id,
            'reading_date' => Carbon::now(),
        ]);
        $this->assertSame($stage2Reading->severity()['text'], 'Stage 2 Hypertension');
    }
}

