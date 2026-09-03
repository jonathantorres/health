<?php

namespace Tests\Browser;

use App\User;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Laravel\Dusk\Browser;
use Tests\DuskTestCase;

class LoginTest extends DuskTestCase
{
    use DatabaseMigrations;

    public function setUp(): void
    {
        parent::setUp();
    }

    /** @test */
    public function unexisting_user_should_not_be_able_to_login()
    {
        $this->browse(function (Browser $browser) {
            $browser->visit('/login')
                    ->assertSee('Login')
                    ->type('email', 'someone@email.com')
                    ->type('password', 'mypass')
                    ->press('Login')
                    ->waitForText('These credentials do not match our records.')
                    ->assertPathIs('/login');
        });
    }

    /** @test */
    public function existing_user_should_be_able_to_login()
    {
        $user = User::factory()->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->visit('/login')
                    ->assertSee('Login')
                    ->type('email', $user->email)
                    ->type('password', 'secret')
                    ->press('Login')
                    ->waitForText('Latest Blood Pressure Readings')
                    ->assertPathIs('/');
        });
    }

    /** @test */
    public function existing_user_should_be_able_to_logout()
    {
        $user = User::factory()->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Blood Pressure Readings')
                    ->click('.navbar-right .dropdown-toggle')
                    ->clickLink('Logout')
                    ->assertPathIs('/login')
                    ->assertSee('Login');
        });
    }
}
